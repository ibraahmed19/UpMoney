import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Deposit, SearchReq } from './deposit.types';
@Injectable({
  providedIn: 'root'
})
export class DepositService {
  private apiUrl = 'http://localhost:8081/api/credits';

  constructor(private _httpClient: HttpClient) { }

  public newDeposit(deposit: Deposit): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', deposit);
  }

  public getDeposits(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getDepositByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchDeposits(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateDeposit(id: any, deposit: Deposit): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, deposit);
  }

  public deleteDeposit(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
