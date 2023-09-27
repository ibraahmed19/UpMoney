import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BankingAccountComponent } from './banking-account.component';

describe('BankingAccountComponent', () => {
  let component: BankingAccountComponent;
  let fixture: ComponentFixture<BankingAccountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BankingAccountComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BankingAccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
