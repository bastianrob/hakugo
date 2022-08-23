import joi, {ValidationError} from "joi";
import jpn from "joi-phone-number";

export const extendedJoi = joi.extend(jpn) as joi.Root;

export const validationErrorToJson = <T = any>(
  initialValue: T,
  error: ValidationError | undefined,
  key?: string,
): T => {
  if (!error) return {} as T;

  return error.details.reduce<T>(
    (acc, entry) => {
      // @ts-ignore
      acc[key || entry.context!.key!] = entry.message;

      return acc;
    },
    {...initialValue},
  );
};
