import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Bill, SearchReq } from './bill.types';

@Injectable({
  providedIn: 'root'
})
export class BillService {
  private apiUrl = 'http://localhost:8081/api/bills';

  constructor(private _httpClient: HttpClient) { }

  public newBill(bill: Bill): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', bill);
  }

  public getBills(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getBillByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchBills(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateBill(id: any, bill: Bill): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, bill);
  }

  public deleteBill(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
