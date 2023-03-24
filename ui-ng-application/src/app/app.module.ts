import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HeaderNavComponent } from './components/header-nav/header-nav.component';
import { MainIntroComponent } from './components/main-intro/main-intro.component';
import { SfiaSelectorFormComponent } from './components/sfia-selector-form/sfia-selector-form.component';
import { SkillSelectorComponent } from './components/skill-selector/skill-selector.component';
import { NgIconsModule } from '@ng-icons/core';
import { bootstrapXCircleFill, bootstrapDashSquare } from '@ng-icons/bootstrap-icons';
import { RoleSelectorComponent } from './components/role-selector/role-selector.component';
import { SelectedRoleViewerComponent } from './components/selected-role-viewer/selected-role-viewer.component';
import { PdpViewerComponent } from './components/pdp-viewer/pdp-viewer.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderNavComponent,
    MainIntroComponent,
    SfiaSelectorFormComponent,
    SkillSelectorComponent,
    RoleSelectorComponent,
    SelectedRoleViewerComponent,
    PdpViewerComponent
  ],
  imports: [
    BrowserModule,
    NgIconsModule.withIcons({ bootstrapXCircleFill, bootstrapDashSquare }),
    AppRoutingModule,
    NgbModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
