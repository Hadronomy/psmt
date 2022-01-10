import { exec, execSync } from 'child_process';
import { Listr } from 'listr2';

interface Ctx { }

const tasks = new Listr<Ctx>(
  [
    {
      title: 'Run Tests',
      task: (ctx) => {
        execSync("jest");
      }
    }
  ],
  {
    concurrent: false
  }
);

tasks.run();
