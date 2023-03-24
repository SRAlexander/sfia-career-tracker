import { TestBed } from '@angular/core/testing';

import { PdpGeneratorService } from './pdp-generator.service';

describe('PdpGeneratorService', () => {
  let service: PdpGeneratorService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PdpGeneratorService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
