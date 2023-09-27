import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { InsuranceProduct, SearchReq } from './insurance-product.types';

@Injectable({
  providedIn: 'root'
})
export class InsuranceProductService {
  private apiUrl = 'http://localhost:8081/api/app/insuranceproduct';

  constructor(private _httpClient: HttpClient) { }

  public NewInsuranceProduct(insuranceProduct: InsuranceProduct): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', insuranceProduct);
  }

  public getInsuranceProducts(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public GetInsuranceProductByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public SearchInsuranceProducts(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public UpdateInsuranceProduct(id: any, insuranceProduct: InsuranceProduct): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, insuranceProduct);
  }

  public DeleteInsuranceProduct(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }

}
