import {ChangeEvent, FormEvent} from "react";

export const formInputChangeToKeyValue =
  <T = string>(callback: (key: string, value: T) => void) =>
  (
    e:
      | FormEvent<HTMLFormElement>
      | ChangeEvent<HTMLInputElement | HTMLAreaElement>,
  ) => {
    const key = (e.target as HTMLInputElement).name;
    const val = (e.target as HTMLInputElement).value;

    callback(key, val as unknown as T);
  };

export const formToJson = <T = any>(form: HTMLFormElement): T => {
  const fd = new FormData(form);
  return Object.fromEntries(fd) as unknown as T;
};

export const formInputChangeToJson =
  <T = any>(callback: (data: T) => void) =>
  (
    e:
      | FormEvent<HTMLFormElement>
      | ChangeEvent<HTMLInputElement | HTMLAreaElement>,
  ) => {
    const form = e.currentTarget as HTMLFormElement;
    const obj = formToJson<T>(form);

    callback(obj);
  };

export const formSubmitToJson =
  <T = any>(callback: (data: T) => void) =>
  (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const obj = formToJson<T>(e.target as HTMLFormElement);

    callback(obj);
  };
