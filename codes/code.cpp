#include <iostream>
#include <vector>
#include <set>
#include <algorithm>

using namespace std;

class Solution {
public:
    int ret;

    void generateSubsets(int index, vector<vector<int>>& ans, vector<int>& nums, vector<int>& subset, int k, int sm, set<vector<int>>& uniqueSubsets) {
        if (index == nums.size()) {
            if (uniqueSubsets.find(subset) == uniqueSubsets.end()) {  
                uniqueSubsets.insert(subset);
                ans.push_back(subset);
                if (sm == k) {
                    ret++;
                }
            }
            return;
        }

        // Include current element
        subset.push_back(nums[index]);
        generateSubsets(index + 1, ans, nums, subset, k, sm + nums[index], uniqueSubsets);

        // Exclude current element
        subset.pop_back();
        generateSubsets(index + 1, ans, nums, subset, k, sm, uniqueSubsets);
    }

    int subsets(vector<int>& arr, int k) {
        ret = 0;
        vector<vector<int>> ans;
        vector<int> subset;
        set<vector<int>> uniqueSubsets;
        sort(arr.begin(), arr.end());
        generateSubsets(0, ans, arr, subset, k, 0, uniqueSubsets);
        return ret;
    }
};

int main() {
    int t;
    cin >> t;
    while (t--) {
        int n;
        cin >> n;
        vector<int> arr(n);
        for (int i = 0; i < n; i++) {
            cin >> arr[i];
        }
        int k;
        cin >> k;
        Solution sol;
        int res = sol.subsets(arr, k);
        cout << res << endl;
    }
    return 0;
}
