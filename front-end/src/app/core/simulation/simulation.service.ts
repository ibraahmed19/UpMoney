import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { SearchReq, Simulation } from './simulation.types';
@Injectable({
  providedIn: 'root'
})
export class SimulationService {
  private apiUrl = 'http://localhost:8081/api/simulations';

  constructor(private _httpClient: HttpClient) { }

  public newSimulation(simulation: Simulation): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/new', simulation);
  }

  public getSimulations(): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/all');
  }

  public getSimulationByID(id: any): Observable<any> {
    return this._httpClient.get(this.apiUrl + '/' + id);
  }

  public searchSimulations(searchReq: SearchReq): Observable<any> {
    return this._httpClient.post(this.apiUrl + '/search', searchReq);
  }

  public updateSimulation(id: any, simulation: Simulation): Observable<any> {
    return this._httpClient.put(this.apiUrl + '/' + id, simulation);
  }

  public deleteSimulation(id: any): Observable<any> {
    return this._httpClient.delete(this.apiUrl + '/' + id);
  }
}
