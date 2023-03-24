import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SfiaSelectorFormComponent } from './sfia-selector-form.component';

describe('SfiaSelectorFormComponent', () => {
  let component: SfiaSelectorFormComponent;
  let fixture: ComponentFixture<SfiaSelectorFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SfiaSelectorFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SfiaSelectorFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
