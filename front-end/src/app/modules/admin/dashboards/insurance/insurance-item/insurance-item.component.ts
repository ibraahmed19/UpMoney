import { Component, Input } from '@angular/core';
import { ProductType } from '../insurance.types';
import { MatDialog } from '@angular/material/dialog';
import { InsuranceProduct } from 'app/core/insurance-product/insurance-product.types';
import { InsuranceDetailsComponent } from '../insurance-details/insurance-details.component';

@Component({
  selector: 'insurance-item',
  templateUrl: './insurance-item.component.html',
  styleUrls: ['./insurance-item.component.scss']
})
export class InsuranceItemComponent {

  @Input() insuranceProduct: InsuranceProduct;

  constructor(
    private _matDialog: MatDialog
  ) {
  }
  openInsuranceProductDetailsDialog(insuranceProduct: InsuranceProduct): void {
    this._matDialog.open(InsuranceDetailsComponent, {
      autoFocus: false,
      data: insuranceProduct
    });
  }
}
