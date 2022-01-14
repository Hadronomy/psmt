import { exec, execSync } from 'child_process';
import { clear } from 'console';
import { Listr } from 'listr2';
import path from 'path';

interface Ctx { }

const tasks = new Listr<Ctx>(
  [
    {
      title: 'Run Tests',
      task: async (ctx) => {
        exec('yarn jest');
      }
    },
    {
      title: 'Lint',
      task: async (ctx) => {
        exec(`yarn eslint --config .eslintrc.json ${path.resolve(__dirname, '..', '..',)}/src/**/*.ts`);
      }
    }
  ],
  {
    concurrent: false
  }
);

tasks.run();
