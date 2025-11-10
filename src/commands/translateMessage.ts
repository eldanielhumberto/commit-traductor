import type { Result } from 'meow';
import { processCommitMessage } from '../utils/processCommitMessage';

export const translateMessage = async (cli: Result<any>) => {
	const file = cli.input[0];

	if (!file) {
		console.error('The commit message file was not provided.');
		process.exit(1);
	}

	processCommitMessage(file);
};
