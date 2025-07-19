import { ColorType } from '@/type/common/ColorType';

export interface TagProps {
  text: string;
  isClose?: boolean;
  color?: ColorType;
  classes?: string[];
  size?: 'small' | 'large';
  isHiddenHash?: boolean;
}
