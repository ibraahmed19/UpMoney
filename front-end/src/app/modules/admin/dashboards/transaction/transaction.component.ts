import { AfterViewInit, ChangeDetectionStrategy, Component, OnDestroy, OnInit, ViewChild, ViewEncapsulation } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Subject, takeUntil } from 'rxjs';
import { MatSort } from '@angular/material/sort';
import { TransactionService } from 'app/core/transaction/transaction.service';
import { MatDialog } from '@angular/material/dialog';
import { TransactionDetailsDialogComponent } from './transaction-details-dialog/transaction-details-dialog.component'
import { Transaction } from 'app/core/transaction/transaction.types';
import { NewTransactionFormComponent } from './new-transaction-form/new-transaction-form.component';

@Component({
  selector: 'app-transaction',
  templateUrl: './transaction.component.html',
  encapsulation: ViewEncapsulation.None,
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class TransactionComponent implements OnInit, AfterViewInit, OnDestroy {
  @ViewChild('recentTransactionsTable', { read: MatSort }) recentTransactionsTableMatSort: MatSort;

  recentTransactionsDataSource: MatTableDataSource<any> = new MatTableDataSource();
  recentTransactionsTableColumns: string[] = ['id', 'amount', 'type', 'execution_date', 'status'];
  private _unsubscribeAll: Subject<any> = new Subject<any>();

  constructor(private _transactionService: TransactionService, private _dialog: MatDialog) { }

  ngOnInit(): void {
    this.reloadTransactions();
  }

  openTransactionDetailsDialog(transaction: Transaction): void {
    this._dialog.open(TransactionDetailsDialogComponent, {
      data: {
        iban_sender: transaction.iban_sender,
        iban_receiver: transaction.iban_receiver,
        description: transaction.description
      }
    });
  }

  openAddTransactionDialog() {
    const dialogRef = this._dialog.open(NewTransactionFormComponent, {
      width: '500px',
      disableClose: true,
    });

    dialogRef.componentInstance.addTransaction.subscribe((transaction: Transaction) => {
      this._transactionService.newTransaction(transaction).subscribe((response) => {
        console.log('Transaction added:', response);
        dialogRef.close();
        this.reloadTransactions();
      });
    });

    dialogRef.componentInstance.cancelForm.subscribe(() => {
      dialogRef.close();
    });
  }

  ngAfterViewInit(): void {
    this.recentTransactionsDataSource.sort = this.recentTransactionsTableMatSort;
  }

  filterByQuery(query: string): void {
    // Implement your filtering logic here
  }

  ngOnDestroy(): void {
    this._unsubscribeAll.next;
    this._unsubscribeAll.complete();
  }

  trackByFn(index: number, item: any): any {
    return item.id || index;
  }

  private reloadTransactions(): void {
    this._transactionService.getTransactionsByBankingAccountId(1)
      .pipe(takeUntil(this._unsubscribeAll))
      .subscribe((data) => {
        this.recentTransactionsDataSource.data = data;
      });
  }
}
