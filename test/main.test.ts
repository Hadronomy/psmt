/**
 * @brief Main program unit tests
 */
import console from 'console';
import { ChildProcessWithoutNullStreams, spawn } from 'child_process';
import path from 'path';

import logger from '../src/services/logger';
import { initCli } from '../src/cli';

describe('CLI', () => {
  const init = jest.fn(initCli);
  beforeEach(() => {
    jest.mock('process');
    jest.spyOn(console, 'log').mockImplementation();
    jest.spyOn(process, 'exit').mockImplementation();
  });
  afterEach(() => {
    jest.unmock('process');
    jest.unmock('console');
  });
  it('Should run', async () => {
    process.argv = ['--help'];
    await init();
    expect(process.exit).toHaveBeenCalled();
    expect(process.exit).not.toHaveBeenCalledWith(1);
  });
  it('Should init', async () => {
    expect(init()).resolves.toBeTruthy();
  });
});

describe('Services', () => {
  describe('API', () => {
    describe('Github', () => {
      it('Is an example', () => {
        expect(true);
      });
    });
  });
});
