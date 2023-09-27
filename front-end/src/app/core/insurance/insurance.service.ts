import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Insurance, SearchReq } from './insurance.types';
@Injectable({
  providedIn: 'root'
})
export class InsuranceService {
  private apiUrl = 'http://localhost:8081/api/insurances';

  constructor(private _httpClient: HttpClient) { }

  public newInsurance(insurance: Insurance): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', insurance);
  }

  public getInsurances(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getInsuranceByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchInsurances(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateInsurance(id: any, insurance: Insurance): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, insurance);
  }

  public deleteInsurance(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
