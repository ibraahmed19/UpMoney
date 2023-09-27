import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Credit, SearchReq } from './credit.types';

@Injectable({
  providedIn: 'root'
})
export class CreditService {
  private apiUrl = 'http://localhost:8081/api/app/credit';

  constructor(private _httpClient: HttpClient) { }

  public newCredit(credit: Credit): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', credit);
  }

  public getCredits(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getCreditByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchCredits(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateCredit(id: any, credit: Credit): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, credit);
  }

  public deleteCredit(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
