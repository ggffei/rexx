
#ifndef HEADER_KRB5_ASN_H
# define HEADER_KRB5_ASN_H

/*
 * #include <krb5.h>
 */
# include <openssl/safestack.h>

#ifdef  __cplusplus
extern "C" {
#endif

/*
 * ASN.1 from Kerberos RFC 1510
 */

/*-     EncryptedData ::=   SEQUENCE {
 *              etype[0]                      INTEGER, -- EncryptionType
 *              kvno[1]                       INTEGER OPTIONAL,
 *              cipher[2]                     OCTET STRING -- ciphertext
 *      }
 */
typedef struct krb5_encdata_st {
    ASN1_INTEGER *etype;
    ASN1_INTEGER *kvno;
    ASN1_OCTET_STRING *cipher;
} KRB5_ENCDATA;

DECLARE_STACK_OF(KRB5_ENCDATA)

/*-     PrincipalName ::=   SEQUENCE {
 *              name-type[0]                  INTEGER,
 *              name-string[1]                SEQUENCE OF GeneralString
 *      }
 */
typedef struct krb5_princname_st {
    ASN1_INTEGER *nametype;
    STACK_OF(ASN1_GENERALSTRING) *namestring;
} KRB5_PRINCNAME;

DECLARE_STACK_OF(KRB5_PRINCNAME)

/*-     Ticket ::=      [APPLICATION 1] SEQUENCE {
 *              tkt-vno[0]                    INTEGER,
 *              realm[1]                      Realm,
 *              sname[2]                      PrincipalName,
 *              enc-part[3]                   EncryptedData
 *      }
 */
typedef struct krb5_tktbody_st {
    ASN1_INTEGER *tktvno;
    ASN1_GENERALSTRING *realm;
    KRB5_PRINCNAME *sname;
    KRB5_ENCDATA *encdata;
} KRB5_TKTBODY;

typedef STACK_OF(KRB5_TKTBODY) KRB5_TICKET;
DECLARE_STACK_OF(KRB5_TKTBODY)

/*-     AP-REQ ::=      [APPLICATION 14] SEQUENCE {
 *              pvno[0]                       INTEGER,
 *              msg-type[1]                   INTEGER,
 *              ap-options[2]                 APOptions,
 *              ticket[3]                     Ticket,
 *              authenticator[4]              EncryptedData
 *      }
 *
 *      APOptions ::=   BIT STRING {
 *              reserved(0), use-session-key(1), mutual-required(2) }
 */
typedef struct krb5_ap_req_st {
    ASN1_INTEGER *pvno;
    ASN1_INTEGER *msgtype;
    ASN1_BIT_STRING *apoptions;
    KRB5_TICKET *ticket;
    KRB5_ENCDATA *authenticator;
} KRB5_APREQBODY;

typedef STACK_OF(KRB5_APREQBODY) KRB5_APREQ;
DECLARE_STACK_OF(KRB5_APREQBODY)

/*      Authenticator Stuff     */

/*-     Checksum ::=   SEQUENCE {
 *              cksumtype[0]                  INTEGER,
 *              checksum[1]                   OCTET STRING
 *      }
 */
typedef struct krb5_checksum_st {
    ASN1_INTEGER *ctype;
    ASN1_OCTET_STRING *checksum;
} KRB5_CHECKSUM;

DECLARE_STACK_OF(KRB5_CHECKSUM)

/*-     EncryptionKey ::=   SEQUENCE {
 *              keytype[0]                    INTEGER,
 *              keyvalue[1]                   OCTET STRING
 *      }
 */
typedef struct krb5_encryptionkey_st {
    ASN1_INTEGER *ktype;
    ASN1_OCTET_STRING *keyvalue;
} KRB5_ENCKEY;

DECLARE_STACK_OF(KRB5_ENCKEY)

/*-     AuthorizationData ::=   SEQUENCE OF SEQUENCE {
 *              ad-type[0]                    INTEGER,
 *              ad-data[1]                    OCTET STRING
 *      }
 */
typedef struct krb5_authorization_st {
    ASN1_INTEGER *adtype;
    ASN1_OCTET_STRING *addata;
} KRB5_AUTHDATA;

DECLARE_STACK_OF(KRB5_AUTHDATA)

/*-     -- Unencrypted authenticator
 *      Authenticator ::=    [APPLICATION 2] SEQUENCE    {
 *              authenticator-vno[0]          INTEGER,
 *              crealm[1]                     Realm,
 *              cname[2]                      PrincipalName,
 *              cksum[3]                      Checksum OPTIONAL,
 *              cusec[4]                      INTEGER,
 *              ctime[5]                      KerberosTime,
 *              subkey[6]                     EncryptionKey OPTIONAL,
 *              seq-number[7]                 INTEGER OPTIONAL,
 *              authorization-data[8]         AuthorizationData OPTIONAL
 *      }
 */
typedef struct krb5_authenticator_st {
    ASN1_INTEGER *avno;
    ASN1_GENERALSTRING *crealm;
    KRB5_PRINCNAME *cname;
    KRB5_CHECKSUM *cksum;
    ASN1_INTEGER *cusec;
    ASN1_GENERALIZEDTIME *ctime;
    KRB5_ENCKEY *subkey;
    ASN1_INTEGER *seqnum;
    KRB5_AUTHDATA *authorization;
} KRB5_AUTHENTBODY;

typedef STACK_OF(KRB5_AUTHENTBODY) KRB5_AUTHENT;
DECLARE_STACK_OF(KRB5_AUTHENTBODY)

/*-  DECLARE_ASN1_FUNCTIONS(type) = DECLARE_ASN1_FUNCTIONS_name(type, type) =
 *      type *name##_new(void);
 *      void name##_free(type *a);
 *      DECLARE_ASN1_ENCODE_FUNCTIONS(type, name, name) =
 *       DECLARE_ASN1_ENCODE_FUNCTIONS(type, itname, name) =
 *        type *d2i_##name(type **a, const unsigned char **in, long len);
 *        int i2d_##name(type *a, unsigned char **out);
 *        DECLARE_ASN1_ITEM(itname) = OPENSSL_EXTERN const ASN1_ITEM itname##_it
 */

DECLARE_ASN1_FUNCTIONS(KRB5_ENCDATA)
DECLARE_ASN1_FUNCTIONS(KRB5_PRINCNAME)
DECLARE_ASN1_FUNCTIONS(KRB5_TKTBODY)
DECLARE_ASN1_FUNCTIONS(KRB5_APREQBODY)
DECLARE_ASN1_FUNCTIONS(KRB5_TICKET)
DECLARE_ASN1_FUNCTIONS(KRB5_APREQ)

DECLARE_ASN1_FUNCTIONS(KRB5_CHECKSUM)
DECLARE_ASN1_FUNCTIONS(KRB5_ENCKEY)
DECLARE_ASN1_FUNCTIONS(KRB5_AUTHDATA)
DECLARE_ASN1_FUNCTIONS(KRB5_AUTHENTBODY)
DECLARE_ASN1_FUNCTIONS(KRB5_AUTHENT)

/* BEGIN ERROR CODES */
/*
 * The following lines are auto generated by the script mkerr.pl. Any changes
 * made after this point may be overwritten when the script is next run.
 */

#ifdef  __cplusplus
}
#endif
#endif
