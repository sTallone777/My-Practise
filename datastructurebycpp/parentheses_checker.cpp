#include<iostream>
using namespace std;

#define MAX 10

class Stack{
    int stk[MAX];

    public:
        int top;
        Stack():top(-1){}
        void push(char);
        char pop();
        void display();
};

void Stack::push(char v){
    if(top == MAX){
        cout<<"The Stack is full \n";
    }else{
        top++;
        stk[top] = v;
    }
}

char Stack::pop(){
    if(top == -1){
        cout<<"The Stack is empty !\n";
        return -1;
    }else{
        char v = stk[top];
        top--;
        return v;        
    }
}

int main(){
    Stack s;
    string exp;
    char tmp;
    int flag = 1;

    cout<<"Enter the expression: ";
    getline(cin, exp);

    for(int i=0; i<exp.length(); i++){
        if(exp[i] == '(' || exp[i] == '{' || exp[i] == '['){
            s.push(exp[i]);
        }
        if(exp[i] == ')' || exp[i] == '}' || exp[i] == ']'){
            if(s.top < 0){
                flag = 0;
            }else{
                tmp = s.pop();
                if((exp[i] == ')' || exp[i] == '}') && tmp == '['){
                    flag = 0;
                }
                if((exp[i] == ']' || exp[i] == '}') && tmp == '('){
                    flag = 0;
                }
                if((exp[i] == ')' || exp[i] == ']') && tmp == '{'){
                    flag = 0;
                }
            }
        }
    }

    if(s.top >= 0){
        flag = 0;
    }

    if(flag == 1){
        cout<<"The expression is valid.\n";
    }else{
        cout<<"The expression is invalid.\n";
    }
    system("pause");
}