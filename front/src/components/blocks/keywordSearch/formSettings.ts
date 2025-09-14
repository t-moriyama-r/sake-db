import { string } from 'yup';

export const validationSchema = {
  keyword: string().required(),
};

export interface FormValues {
  keyword: string;
}

export const onSubmitCore = async () => {
  return true;
};
