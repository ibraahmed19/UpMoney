import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreditDetailsDialogComponent } from './credit-details-dialog.component';

describe('CreditDetailsDialogComponent', () => {
  let component: CreditDetailsDialogComponent;
  let fixture: ComponentFixture<CreditDetailsDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreditDetailsDialogComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreditDetailsDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
