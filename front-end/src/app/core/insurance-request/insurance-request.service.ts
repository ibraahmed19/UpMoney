import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { InsuranceRequest, SearchReq } from './insurance-request.types';
@Injectable({
  providedIn: 'root'
})
export class InsuranceRequestService {
  private apiUrl = 'http://localhost:8081/api/insurancerequests';

  constructor(private _httpClient: HttpClient) { }

  public newInsuranceRequest(insuranceRequest: InsuranceRequest): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', insuranceRequest);
  }

  public getInsuranceRequests(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getInsuranceRequestByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchInsuranceRequests(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateInsuranceRequest(id: any, insuranceRequest: InsuranceRequest): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, insuranceRequest);
  }

  public deleteInsuranceRequest(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
