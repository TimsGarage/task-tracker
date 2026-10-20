const API_BASE_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080/api';

export class ApiError extends Error {
	constructor(message, status) {
		super(message);
		this.name = 'ApiError';
		this.status = status;
	}
}

export async function apiRequest(path, options = {}) {
	const response = await fetch(`${API_BASE_URL}${path}`, {
		...options,
		headers: {
			...(options.body ? { 'Content-Type': 'application/json' } : {}),
			...options.headers
		}
	});

	if (!response.ok) {
		const message = (await response.text()) || 'Request failed';
		throw new ApiError(message, response.status);
	}

	if (response.status === 204) {
		return null;
	}

	return response.json();
}
