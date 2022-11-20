# cazilla
An alternative CA list solution. Inspired by [SagerNet](https://github.com/SagerNet/SagerNet).

---

System CA list is not always reliable for some cases, while users/manufacturers could add their self-signed certificate to CA list and do Man-in-the-middle attacks easily.  
Cazilla is based on [Mozilla CA Certificate Program](https://wiki.mozilla.org/CA/Included_Certificates), providing methods to easily download Mozilla included CA list and apply them to your code.  
  
## Examples

See [examples](https://github.com/H1JK/cazilla/tree/master/examples).

## Credits

This includes Included CA Certificates (CSV with PEM of raw certificate data) provided by Mozilla.  
Users must agree [CCADB Data Usage Terms](https://www.ccadb.org/rootstores/usage#ccadb-data-usage-terms) before use.  
Code parts are licensed under [MIT License](https://github.com/H1JK/cazilla/blob/master/LICENSE).