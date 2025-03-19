import java.util.*;

class Solution {
    private int ret;

    private void generateSubsets(int index, List<List<Integer>> ans, int[] nums, List<Integer> subset, int k, int sm, Set<List<Integer>> uniqueSubsets) {
        if (index == nums.length) {
            if (!uniqueSubsets.contains(subset)) {
                uniqueSubsets.add(new ArrayList<>(subset));
                ans.add(new ArrayList<>(subset));
                if (sm == k) {
                    ret++;
                }
            }
            return;
        }

        // Include current element
        subset.add(nums[index]);
        generateSubsets(index + 1, ans, nums, subset, k, sm + nums[index], uniqueSubsets);

        // Exclude current element
        subset.remove(subset.size() - 1);
        generateSubsets(index + 1, ans, nums, subset, k, sm, uniqueSubsets);
    }

    public int subsets(int[] arr, int k) {
        ret = 0;
        List<List<Integer>> ans = new ArrayList<>();
        List<Integer> subset = new ArrayList<>();
        Set<List<Integer>> uniqueSubsets = new HashSet<>();
        Arrays.sort(arr);
        generateSubsets(0, ans, arr, subset, k, 0, uniqueSubsets);
        return ret;
    }
}

public class Main {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int t = scanner.nextInt();
        while (t-- > 0) {
            int n = scanner.nextInt();
            int[] arr = new int[n];
            for (int i = 0; i < n; i++) {
                arr[i] = scanner.nextInt();
            }
            int k = scanner.nextInt();
            Solution sol = new Solution();
            int res = sol.subsets(arr, k);
            System.out.println(res);
        }
        scanner.close();
    }
}
