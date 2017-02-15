#pragma once

#ifndef NULL
#define NULL 0
#endif

#include <iostream>

namespace ZY
{
	class string
	{
	public:
		string(const char *str = NULL);
		string(const string &str); // copy constructor
		~string();

	public:
		// assign constructor
		string& operator=(const string& str);

		// overload operator []
		char operator[](int nIdx) const;

		// overload IO
		friend std::ostream& operator<<(std::ostream& out, const string &str);
		friend std::istream& operator>>(std::istream& in, const string &str);

		// return length of the string
		int length() const;
	private:
		char *m_pStr;
	};
}