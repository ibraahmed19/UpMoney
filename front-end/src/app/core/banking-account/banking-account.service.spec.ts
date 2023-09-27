import { TestBed } from '@angular/core/testing';

import { BankingAccountService } from './banking-account.service';

describe('BankingAccountService', () => {
  let service: BankingAccountService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BankingAccountService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
