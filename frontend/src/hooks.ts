import cookie from 'cookie';
import { v4 as uuid } from '@lukeed/uuid';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const request = event.request;
	// const cookies = cookie.parse(request.headers.cookie || '');
	// event.locals["userid"]
	// request.locals.userid = uuid();

	// // TODO https://github.com/sveltejs/kit/issues/1046
	// if (request.query.has('_method')) {
	// 	request.method = request.query.get('_method').toUpperCase();
	// }

	const response = await resolve(event, {
		ssr: false,
	});

	// if (!cookies.userid) {
	// 	// if this is the first time the user has visited this app,
	// 	// set a cookie so that we recognise them when they return
	// 	response.headers['set-cookie'] = cookie.serialize('userid', request.locals.userid, {
	// 		path: '/',
	// 		httpOnly: true
	// 	});
	// }

	return response;
};
