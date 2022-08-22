import {ChangeEvent, FormEvent} from "react";

export const formInputChangeToJson =
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

export const formSubmitToJson =
  <T = any>(callback: (data: T) => void) =>
  (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const fd = new FormData(e.target as HTMLFormElement);

    callback(Object.fromEntries(fd) as unknown as T);
  };
