import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Paymee, searchReq } from './paymee.types';
@Injectable({
  providedIn: 'root'
})
export class PaymeeService {
  private apiUrl = 'http://localhost:8081/api/payments';

  constructor(private _httpClient: HttpClient) { }

  public newPayment(paymee: Paymee): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', paymee);
  }

  public getPayments(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getPaymentByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchPayments(searchReq: searchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }
}
