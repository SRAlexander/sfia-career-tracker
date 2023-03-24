import { TestBed } from '@angular/core/testing';

import { SfiaSkillsService } from './sfia-skills.service';

describe('SfiaSkillsService', () => {
  let service: SfiaSkillsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SfiaSkillsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
