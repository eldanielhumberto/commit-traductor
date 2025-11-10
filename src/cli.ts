#!/usr/bin/env bun
import meow from 'meow';
import { translateMessage } from './commands/translateMessage';

const cli = meow(
	`
  🧠 commit-traductor`,
	{
		importMeta: import.meta,
	}
);

translateMessage(cli);
