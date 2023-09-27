import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { CreditRequest, SearchReq } from './credit-request.types';
@Injectable({
  providedIn: 'root'
})
export class CreditRequestService {
  private apiUrl = 'http://localhost:8081/api/creditrequests';

  constructor(private _httpClient: HttpClient) { }

  public newCreditRequest(creditRequest: CreditRequest): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', creditRequest);
  }

  public getCreditRequests(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getCreditRequestByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchCreditRequests(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateCreditRequest(id: any, creditRequest: CreditRequest): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, creditRequest);
  }

  public deleteCreditRequest(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
