#ifndef _HTTP_REQUEST_H_
#define _HTTP_REQUEST_H_

#include<string>
#include<vector>
#include<fstream>
#include<Http.h>

using namespace std;

class HttpRequest{

public:
    HttpRequest();
    ~HttpRequest();

    void printRequest(void );
    void addData(const char *, const int& );
    void addRequestBody(const string& );

    int setMethod(Method );
    Method getMethod(void );

    int setURL(string );
    string getURL(void );

    int setProtocol(Protocol );
    Protocol getProtocol(void );

    int setUserAgent(string userAgent);
    string getUserAgent(void );

    int setHttpHeader(string headerName, string headerContent);
    string getHttpHeader(string headerName);

    int setHttpHeaderVector(vector<pair<string, string>>* const);
    vector<pair<string, string>>* getHttpHeaderVector(void );

    int setRequestBody(const string*);
    string* getRequestBodyPtr(void );

    int parseRequest(void);
    int prepareRequest(void );
    size_t getRequestSize(void );
    string* getRequestDataPtr(void );

    int copyToFile(ofstream&);
    int conpyFromFile(ifstream&, size_t);

private:
    Method m_httpMethod;
    string m_url;
    Protocol m_protocol;
    string m_hostName;
    string m_userAgent;

    vector<pair<string, string>> m_httpHeaders;
    string m_requestBody;
    string m_data;
};

#endif