# Four letter abbreviations

```
/*!
	@enum CertificateItemAttributes
	@abstract Indicates the type of a certificate item attribute.
	@constant kSecSubjectItemAttr Indicates a DER-encoded subject distinguished name.
	@constant kSecIssuerItemAttr Indicates a DER-encoded issuer distinguished name.
	@constant kSecSerialNumberItemAttr Indicates a DER-encoded certificate serial number (without the tag and length).
	@constant kSecPublicKeyHashItemAttr Indicates a public key hash.
	@constant kSecSubjectKeyIdentifierItemAttr Indicates a subject key identifier.
	@constant kSecCertTypeItemAttr Indicates a certificate type.
	@constant kSecCertEncodingItemAttr Indicates a certificate encoding.
*/
enum
{
    kSecSubjectItemAttr 			 = 'subj',
    kSecIssuerItemAttr 				 = 'issu',
    kSecSerialNumberItemAttr     	 = 'snbr',
    kSecPublicKeyHashItemAttr    	 = 'hpky',
    kSecSubjectKeyIdentifierItemAttr = 'skid',
	kSecCertTypeItemAttr		 	 = 'ctyp',
	kSecCertEncodingItemAttr	 	 = 'cenc'
};
```
