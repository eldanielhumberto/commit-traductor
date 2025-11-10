import translate from 'translate';

export const renameCommit = async (msg: string): Promise<string> => {
	translate.engine = 'google';
	return await translate(msg, { from: 'es', to: 'en' });
};
