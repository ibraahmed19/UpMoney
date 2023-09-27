import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Bank, SearchReq } from './bank.types';

@Injectable({
  providedIn: 'root'
})
export class BankService {
  private apiUrl = 'http://localhost:8081/api/banks';

  constructor(private _httpClient: HttpClient) { }

  public NewBank(bank: Bank): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', bank);
  }

  public GetBanks(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public GetBankByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public SearchBanks(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public UpdateBank(id: any, bank: Bank): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, bank);
  }

  public DeleteBank(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }

}
