#ifndef	__LIBCOMMON_H__
#define	__LIBCOMMON_H__


void spi_init(void);
void spiDataWrite(unsigned char data);

unsigned char spiDataRead(unsigned char data);

#endif