import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PdpViewerComponent } from './pdp-viewer.component';

describe('PdpViewerComponent', () => {
  let component: PdpViewerComponent;
  let fixture: ComponentFixture<PdpViewerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PdpViewerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PdpViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
