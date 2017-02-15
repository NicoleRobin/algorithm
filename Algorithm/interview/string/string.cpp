#include "string.h"

#include <cstring>
using namespace std;

namespace ZY
{
	string::string(const char *s /*= NULL*/)
	{
			int nLen = strlen(s);
			m_pStr = new char[nLen + 1];
			memset(m_pStr, 0, nLen + 1);
			strcpy(m_pStr, s);
	}

	string::string(const string &str)
	{
		m_pStr = new char[str.length() + 1];
		memset(m_pStr, 0, str.length() + 1);
		strcpy(m_pStr, str.m_pStr);
	}

	string::~string()
	{
		if (m_pStr != NULL)
		{
			delete[] m_pStr;
			m_pStr = NULL;
		}
	}

	string& string::operator=(const string &str)
	{
		if (this != &str)
		{
			delete m_pStr;
			m_pStr = new char[str.length() + 1];
			memset(m_pStr, 0, str.length() + 1);
			strcpy(m_pStr, str.m_pStr);
		}

		return *this;
	}

	char string::operator[](int nIdx) const
	{
		return m_pStr[nIdx];
	}

	ostream& operator<<(ostream& out, const string &str)
	{
		int nLen = str.length();
		for (int i = 0; i < nLen; ++i)
		{
			out << str[i];
		}

		return out;
	}

	istream& operator>>(istream& in, const string &str)
	{
		return in;
	}

	int string::length() const
	{
		return strlen(m_pStr);
	}
}