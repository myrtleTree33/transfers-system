import { browser } from '$app/environment';
import type { Writable } from 'svelte/store';
import { writable, get } from 'svelte/store';

// Persists store to local storage
// See https://stackoverflow.com/questions/56488202/how-to-persist-svelte-store
const storage = <T>(key: string, initValue: T): Writable<T> => {
  const store = writable(initValue);
  if (!browser) return store;

  const storedValueStr = localStorage.getItem(key);
  if (storedValueStr != null) store.set(JSON.parse(storedValueStr));

  store.subscribe((val) => {
    if ([null, undefined].includes(val)) {
      localStorage.removeItem(key);
    } else {
      localStorage.setItem(key, JSON.stringify(val));
    }
  });

  window.addEventListener('storage', () => {
    const storedValueStr = localStorage.getItem(key);
    if (storedValueStr == null) return;

    const localValue: T = JSON.parse(storedValueStr);
    if (localValue !== get(store)) store.set(localValue);
  });

  return store;
};

export default storage;
