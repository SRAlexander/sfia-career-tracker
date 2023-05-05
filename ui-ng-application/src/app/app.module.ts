import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NgSelectModule } from '@ng-select/ng-select';

import { NgIconsModule } from '@ng-icons/core';
import { bootstrapXCircleFill, bootstrapDashSquare, bootstrapMarkdown, bootstrapTable, bootstrapDownload } from '@ng-icons/bootstrap-icons';

import { HeaderNavComponent } from './components/header-nav/header-nav.component';
import { MainIntroComponent } from './components/main-intro/main-intro.component';
import { RoleSelectorComponent } from './components/pdp-generator-container/role-selector-section/role-selector.component';
import { SelectedRoleViewerComponent } from './components/selected-role-viewer/selected-role-viewer.component';
import { PdpGeneratorContainerComponent } from './components/pdp-generator-container/pdp-generator-container.component';
import { SfiaSelectorFormComponent } from './components/pdp-generator-container/sfia-selector-section/sfia-selector-section.component';
import { CreationTypeSectionComponent } from './components/pdp-generator-container/creation-type-section/creation-type-section.component';
import { SkillSelectorComponent } from './components/pdp-generator-container/skill-selector-section/skill-selector.component';
import { CustomTypeaheadComponent } from './components/shared/custom-typeahead/custom-typeahead.component';
import { PdpViewerComponent } from './components/pdp-generator-container/pdp-viewer-section/pdp-viewer.component';
import { FormsModule, NgSelectOption } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    HeaderNavComponent,
    MainIntroComponent,
    SfiaSelectorFormComponent,
    SkillSelectorComponent,
    RoleSelectorComponent,
    SelectedRoleViewerComponent,
    PdpViewerComponent,
    PdpGeneratorContainerComponent,
    CreationTypeSectionComponent,
    RoleSelectorComponent
  ],
  imports: [
    BrowserModule,
    NgIconsModule.withIcons({ bootstrapXCircleFill, bootstrapDashSquare, bootstrapMarkdown, bootstrapTable, bootstrapDownload }),
    AppRoutingModule,
    NgbModule,
    CustomTypeaheadComponent,
    NgSelectModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
