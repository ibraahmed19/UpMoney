import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { BankingAccount, SearchReq } from './banking-account.types';
@Injectable({
  providedIn: 'root'
})
export class BankingAccountService {
  private apiUrl = 'http://localhost:8081/api/bankingaccounts';

  constructor(private _httpClient: HttpClient) { }

  public newBankingAccount(bankingAccount: BankingAccount): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', bankingAccount);
  }

  public getBankingAccounts(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getBankingAccountByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public getBankingAccountsByClientId(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchBankingAccounts(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateBankingAccount(id: any, bankingAccount: BankingAccount): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, bankingAccount);
  }

  public deleteBankingAccount(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
