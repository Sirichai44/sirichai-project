import { extendTheme } from '@mui/joy';

const theme = extendTheme({
  components: {
    JoyButton: {
      defaultProps: {
        size: 'sm'
      },
      styleOverrides: {
        root: {
          borderRadius: '0.375rem'
        }
      }
    },
    JoyInput: {
      defaultProps: {
        size: 'sm'
      },
      styleOverrides: {
        root: {
          borderRadius: '0.375rem'
        }
      }
    },
    JoySelect: {
      defaultProps: {
        size: 'sm'
      },
      styleOverrides: {
        root: {
          borderRadius: '0.375rem'
        }
      }
    },
    JoyTooltip: {
      defaultProps: {
        size: 'sm'
      }
    }
  },
  fontFamily: {
    body: 'geist'
  }
});

export default theme;
