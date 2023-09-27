import { Component, Input, Inject } from '@angular/core';
import { InsuranceProduct } from 'app/core/insurance-product/insurance-product.types';
import { Insurance } from 'app/core/insurance/insurance.types';
import { ProductType } from '../insurance.types';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'app-insurance-details',
  templateUrl: './insurance-details.component.html',
  styleUrls: ['./insurance-details.component.scss']
})
export class InsuranceDetailsComponent {

  insuranceProduct: InsuranceProduct;

  constructor(
    @Inject(MAT_DIALOG_DATA) private  _data: InsuranceProduct ,
    private _matDialogRef: MatDialogRef<InsuranceDetailsComponent>
  ) {
    this.insuranceProduct = _data;
  }


  
}
