import { TestBed } from '@angular/core/testing';

import { InsuranceRequestService } from './insurance-request.service';

describe('InsuranceRequestService', () => {
  let service: InsuranceRequestService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(InsuranceRequestService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
