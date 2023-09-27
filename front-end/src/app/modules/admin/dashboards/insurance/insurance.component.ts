import { Component, OnInit  } from '@angular/core';
import { ProductType } from './insurance.types';
import { InsuranceProductService } from 'app/core/insurance-product/insurance-product.service';
import { InsuranceProduct } from 'app/core/insurance-product/insurance-product.types';

@Component({
  selector: 'app-insurance',
  templateUrl: './insurance.component.html',
  styleUrls: ['./insurance.component.scss']
})
export class InsuranceComponent  implements OnInit {

  insuranceProducts: InsuranceProduct[];
  public product: typeof ProductType = ProductType;


  constructor(private insuranceProductService: InsuranceProductService){


  }
  ngOnInit(): void {
    this.insuranceProductService.getInsuranceProducts().subscribe(res => {
      this.insuranceProducts = res;
    })
  }

}
