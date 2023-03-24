import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SelectedRoleViewerComponent } from './selected-role-viewer.component';

describe('SelectedRoleViewerComponent', () => {
  let component: SelectedRoleViewerComponent;
  let fixture: ComponentFixture<SelectedRoleViewerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SelectedRoleViewerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SelectedRoleViewerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
