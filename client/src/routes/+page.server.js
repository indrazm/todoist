import { SECRET_API_URL } from '$env/static/private';

export async function load() {
	console.log(SECRET_API_URL);
}
