#!/bin/node
import { initCli } from '@/cli';

/**
 * `psmt` main program
 */
async function mainProgram() {
  await initCli(process.argv);
}

mainProgram();
