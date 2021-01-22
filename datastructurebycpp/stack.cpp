#include <iostream>

using namespace std;

#define MAX 10

class Stack
{
    int stk[MAX], top;

public:
    //Stack():top(-1){};
    Stack();
    void push(int);
    int pop();
    void display();
};

Stack::Stack(){
    top = -1;
}

void Stack ::push(int value)
{
    if (top == MAX - 1)
    {
        cout << "The Stack is full !\n";
    }
    else
    {
        top++;
        stk[top] = value;
    }
}

int Stack ::pop()
{
    if (top == -1)
    {
        cout << "The Stack is empty !\n";
        return -1;
    }
    else
    {
        int value = stk[top];

        top--;
        return value;
    }
}

void Stack ::display()
{
    if (top == -1)
    {
        cout << "The Stack is empty !\n";
    }
    else
    {
        for (int i = top; i >= 0; i--)
        {
            cout << stk[i] << " ";
        }
        cout << endl;
    }
}

int main()
{
    int option, value;
    Stack s;

    do
    {
        cout << "\n**** MENU ****\n"
             << "1. Push\n"
             << "2. Pop\n"
             << "3. Display\n"
             << "4. Exit\n";

        cout << "Enter your option: ";
        cin >> option;

        switch (option)
        {
        case 1:
            cout << "Enter the value to be pushed:";
            cin >> value;

            s.push(value);
            break;
        case 2:
            value = s.pop();
            if (value != -1)
            {
                cout << "The popped element is: " << value << endl;
            }
            break;
        case 3:
            s.display();
            break;
        case 4:
            break;
        default:
            cout << "Wrong option !\n";
            break;
        }
    } while (option != 4);

    if (option == 4)
    {
        cout << "\nThank you for using the program !\n"
             << "Have a good day. \n\n";
    }
}