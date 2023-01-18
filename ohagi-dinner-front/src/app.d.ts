/// <reference types="@sveltejs/kit" />
/// <reference types="@sveltejs/adapter-cloudflare-workers" />

declare namespace App {
	// interface Error {}
	// interface Locals {}
	// interface PageData {}
	interface Platform {
	  env: {
		YOUR_KV_NAMESPACE: KVNamespace
		YOUR_DURABLE_OBJECT_NAMESPACE: DurableObjectNamespace
	  }
	}
  }