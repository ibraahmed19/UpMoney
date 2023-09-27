import { AfterViewInit, ChangeDetectionStrategy, Component, OnDestroy, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { CreditService } from 'app/core/credit/credit.service';
import { Credit } from 'app/core/credit/credit.types';
import { ApexOptions } from 'ng-apexcharts';
import { Subject, takeUntil } from 'rxjs';
import { CreditDetailsDialogComponent } from './credit-details-dialog/credit-details-dialog.component';
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-credit',
  templateUrl: './credit.component.html',
  styleUrls: ['./credit.component.scss'],
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class CreditComponent implements OnInit, AfterViewInit, OnDestroy {

  @ViewChild('recentTransactionsTable', {read: MatSort}) recentCreditsTableMatSort: MatSort;
  
  data: any;
  accountBalanceOptions: ApexOptions;
  private _unsubscribeAll: Subject<any> = new Subject<any>();

  recentCreditsDataSource: MatTableDataSource<any> = new MatTableDataSource();
  recentCreditsTableColumns: string[] = ['id', 'amount', 'type', 'duration' ,'status'];
  


  constructor(private _creditService: CreditService, private _dialog: MatDialog){}
  ngOnInit(): void
  {
    this.reloadCredits();
  }

  openCreditDetailsDialog(credit: Credit): void {
    this._dialog.open(CreditDetailsDialogComponent, {
      data: {
        start_date: credit.start_date,
        end_date: credit.end_date,
        monthly_payments: credit.monthly_payments,
        comments: credit.comments
      }
    });
  }
  


  ngAfterViewInit(): void
  {
      // Make the data source sortable
      this.recentCreditsDataSource.sort = this.recentCreditsTableMatSort;
  }
  filterByQuery(query: string): void {
    // Implement your filtering logic here
  }
  /**
   * On destroy
   */
  ngOnDestroy(): void
  {
      // Unsubscribe from all subscriptions
      this._unsubscribeAll.next(null);
      this._unsubscribeAll.complete();
  }

      /**
     * Track by function for ngFor loops
     *
     * @param index
     * @param item
     */
      trackByFn(index: number, item: any): any
      {
          return item.id || index;
      }
      private reloadCredits(): void {
        this._creditService.getCredits()
          .pipe(takeUntil(this._unsubscribeAll))
          .subscribe((data) => {
            this.recentCreditsDataSource.data = data;
          });
      }
  

}

