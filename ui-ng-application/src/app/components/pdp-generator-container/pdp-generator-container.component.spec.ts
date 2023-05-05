import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PdpGeneratorContainerComponent } from './pdp-generator-container.component';

describe('PdpGeneratorContainerComponent', () => {
  let component: PdpGeneratorContainerComponent;
  let fixture: ComponentFixture<PdpGeneratorContainerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PdpGeneratorContainerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PdpGeneratorContainerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
