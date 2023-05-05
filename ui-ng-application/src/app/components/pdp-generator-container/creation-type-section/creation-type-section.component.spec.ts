import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreationTypeSectionComponent } from './creation-type-section.component';

describe('CreationTypeSectionComponent', () => {
  let component: CreationTypeSectionComponent;
  let fixture: ComponentFixture<CreationTypeSectionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreationTypeSectionComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreationTypeSectionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
