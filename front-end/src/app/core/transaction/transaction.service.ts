import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { SearchReq, Transaction } from './transaction.types';

@Injectable({
  providedIn: 'root'
})
export class TransactionService {
  private apiUrl = 'http://localhost:8081/api/app/transaction';

  constructor(private _httpClient: HttpClient) { }


  public getTransactionsByBankingAccountId(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all/' + id);
  }

  public getTransactions(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getTransactionByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchTransactions(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public newTransaction(transaction: Transaction): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', transaction);
  }

  public updateTransaction(id: any, transaction: Transaction): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, transaction);
  }

  public deleteTransaction(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }

}
