import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { InsurancePolicy, SearchReq } from './insurance-policy.types';
@Injectable({
  providedIn: 'root'
})
export class InsurancePolicyService {
  private apiUrl = 'http://localhost:8081/api/insurancepolicies';

  constructor(private _httpClient: HttpClient) { }

  public newInsurancePolicy(insurancePolicy: InsurancePolicy): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', insurancePolicy);
  }

  public getInsurancePolicies(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getInsurancePolicyByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchInsurancePolicies(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateInsurancePolicy(id: any, insurancePolicy: InsurancePolicy): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, insurancePolicy);
  }

  public deleteInsurancePolicy(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
