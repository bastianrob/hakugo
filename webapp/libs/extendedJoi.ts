import joi, {ValidationError} from "joi";
import jpn from "joi-phone-number";

export const extendedJoi = joi.extend(jpn) as joi.Root;

export const validationErrorToJson = <T = any>(
  error: ValidationError | undefined,
): T => {
  if (!error) return {} as T;

  return error.details.reduce<T>((acc: any, entry) => {
    acc[entry.context!.key!] = entry.message;

    return acc;
  }, {} as any);
};
