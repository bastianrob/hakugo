export const getInitialFromEmail = (email: string): string => {
  const address = email.split("@", 2)[0];
  const containsSeparator = new RegExp("[_.-]");
  if (containsSeparator.test(address)) {
    const [first, second] = address.split(containsSeparator, 2);
    return `${first[0]}${second[0]}`.toUpperCase();
  } else {
    return address.substring(0, 2).toUpperCase();
  }
};
